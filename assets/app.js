const ChunksModes = {
    OFF: '',
    ON: 'chunks'
};

const PushModes = {
    OFF: '',
    ON: 'push'
};

const getQueryString = (params = {}) => {
    return [...Object.keys(params)]
        .map(key => `${key}=${params[key]}`)
        .join('&');
};

const app = {
    chunkSize: 150,
    portfolioSize: 100,
    initialLimit: 50,

    accounts: [],
    portfolios: {},

    PortfolioApp: function(service, ctors) {
        this.service = service;
        this.ctors = ctors;

        const { searchParams: params } = (new URL(document.location));

        this.chunksMode = params.get('chunks') || ChunksModes.OFF;
        this.includePortfolio = params.get('includePortfolio');
        this.portfolioSize = Math.min(params.get('size') || this.portfolioSize, 30000);
        this.pushMode = params.get('push') || PushModes.OFF;
    }
};

app.PortfolioApp.prototype.initialize = function(callback) {
    return this.fetchAccounts('', { 'X-Portfolio-Size': this.portfolioSize })
        .then(data => {
            app.accounts = data || [];

			setAccountsDropdown('accounts', app.accounts);

            callback && callback(app.accounts);

            return this.grpc_fetchPortfolio(app.accounts[0]);
        });
};

// ----- REST API Client ----- //

app.PortfolioApp.prototype.fetchAccounts = function(queryString = '', headers = {}) {
    return fetch('/api/v1/accounts' + queryString, { headers })
        .then(r => r.json())
};

app.PortfolioApp.prototype.fetchPortfolio = function({ accountId }, params = {}, headers = {}) {
    const qs = getQueryString(params);

    return fetch(`/api/v1/portfolios/${accountId}${qs && '?' + qs}`, {
        headers: {
            'X-Enable-Chunks': chunksMode,
            'X-Enable-Push': pushMode,
            'X-Portfolio-Size': portfolioSize,
            ...headers
        }
    })
        .then(r => r.json());
};

// ----- gRPC API Client ----- //

app.PortfolioApp.prototype.grpc_fetchPortfolio = function(account) {
    const request = new this.ctors.Account();

    request.setId(account.accountId);
    request.setEncodedId(account.encAccountId);
    request.setName(account.accountLongName);
    request.setShortName(account.accountShortName);
    request.setDescription(account.acctDesc);
    request.setMode(account.accountMode);
    request.setStatus(account.acctStatus);
    request.setType(account.acctType);
    request.setCreated(account.acctCreationDate);

    return this.service.getHoldings(request, { 'X-Portfolio-Size': this.portfolioSize });
};

app.PortfolioApp.prototype.grpc_fetchHoldings = function(account) {
    const request = new this.ctors.Account();

    request.setId(account.accountId);
    request.setEncodedId(account.encAccountId);
    request.setName(account.accountLongName);
    request.setShortName(account.accountShortName);
    request.setDescription(account.acctDesc);
    request.setMode(account.accountMode);
    request.setStatus(account.acctStatus);
    request.setType(account.acctType);
    request.setCreated(account.acctCreationDate);

    return this.service.listHoldings(request);
};

module.exports = app;
