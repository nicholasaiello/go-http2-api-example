const { Account, Holding, Holdings } = require('./js/holdings_pb.js');
const { AccountHoldingsPromiseClient } = require('./js/holdings_grpc_web_pb.js');
const { PortfolioApp } = require('../assets/app.js');

const grpc = {};
grpc.web = require('grpc-web');

const port = document.location.protocol === 'http:' ? 8081 : 8082;
const portfolioService = new AccountHoldingsPromiseClient(
    document.location.protocol + '//changeforabutton.local:' + port, null, null);

const grpcApp = new PortfolioApp(portfolioService, { Account, Holding, Holdings });

grpcApp.initialize()
    .then(portfolio => {
        console.log('grpc_fetchPortfolio.getHoldings', portfolio);

        app.portfolio = portfolio;

		appendPositionsRows(portfolio.getPositionsList(), false, mapProtoPositionRow);
    });
