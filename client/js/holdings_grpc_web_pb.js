/**
 * @fileoverview gRPC-Web generated client stub for proto
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.proto = require('./holdings_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.AccountHoldingsClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.AccountHoldingsPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.Account,
 *   !proto.proto.Holdings>}
 */
const methodDescriptor_AccountHoldings_GetHoldings = new grpc.web.MethodDescriptor(
  '/proto.AccountHoldings/GetHoldings',
  grpc.web.MethodType.UNARY,
  proto.proto.Account,
  proto.proto.Holdings,
  /**
   * @param {!proto.proto.Account} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Holdings.deserializeBinary
);


/**
 * @param {!proto.proto.Account} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.proto.Holdings)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.Holdings>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.AccountHoldingsClient.prototype.getHoldings =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.AccountHoldings/GetHoldings',
      request,
      metadata || {},
      methodDescriptor_AccountHoldings_GetHoldings,
      callback);
};


/**
 * @param {!proto.proto.Account} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.Holdings>}
 *     Promise that resolves to the response
 */
proto.proto.AccountHoldingsPromiseClient.prototype.getHoldings =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.AccountHoldings/GetHoldings',
      request,
      metadata || {},
      methodDescriptor_AccountHoldings_GetHoldings);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.Holding,
 *   !proto.proto.Holding>}
 */
const methodDescriptor_AccountHoldings_GetHolding = new grpc.web.MethodDescriptor(
  '/proto.AccountHoldings/GetHolding',
  grpc.web.MethodType.UNARY,
  proto.proto.Holding,
  proto.proto.Holding,
  /**
   * @param {!proto.proto.Holding} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Holding.deserializeBinary
);


/**
 * @param {!proto.proto.Holding} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.proto.Holding)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.Holding>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.AccountHoldingsClient.prototype.getHolding =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.AccountHoldings/GetHolding',
      request,
      metadata || {},
      methodDescriptor_AccountHoldings_GetHolding,
      callback);
};


/**
 * @param {!proto.proto.Holding} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.Holding>}
 *     Promise that resolves to the response
 */
proto.proto.AccountHoldingsPromiseClient.prototype.getHolding =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.AccountHoldings/GetHolding',
      request,
      metadata || {},
      methodDescriptor_AccountHoldings_GetHolding);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.Account,
 *   !proto.proto.Holding>}
 */
const methodDescriptor_AccountHoldings_ListHoldings = new grpc.web.MethodDescriptor(
  '/proto.AccountHoldings/ListHoldings',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.proto.Account,
  proto.proto.Holding,
  /**
   * @param {!proto.proto.Account} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Holding.deserializeBinary
);


/**
 * @param {!proto.proto.Account} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.proto.Holding>}
 *     The XHR Node Readable Stream
 */
proto.proto.AccountHoldingsClient.prototype.listHoldings =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/proto.AccountHoldings/ListHoldings',
      request,
      metadata || {},
      methodDescriptor_AccountHoldings_ListHoldings);
};


/**
 * @param {!proto.proto.Account} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.proto.Holding>}
 *     The XHR Node Readable Stream
 */
proto.proto.AccountHoldingsPromiseClient.prototype.listHoldings =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/proto.AccountHoldings/ListHoldings',
      request,
      metadata || {},
      methodDescriptor_AccountHoldings_ListHoldings);
};


module.exports = proto.proto;

