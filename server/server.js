const path = require("path");
const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
const products = require("./sample-products");
const assert = require("assert");

const PORT = 30043;

const PROTO_PATH = path.resolve(__dirname, "../proto/products.proto");

// load the JS version of the proto to create a GRPC object
var packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});
var protoDescriptor = grpc.loadPackageDefinition(packageDefinition);

// get the products gRPC package
var productsPackage = protoDescriptor.products;

// define and implement the products service gRPC server
const getGrpcServer = () => {
  const server = new grpc.Server();
  const productsService = productsPackage.ProductsService.service;

  server.addService(productsService, {
    getAllProducts: (_, callback) => {
      console.log("[GRPC Server]: Getting all products...");
      callback(null, { products });
    },

    getProductsInCategory: (call, callback) => {
      var categoryName = call.request.name;

      console.log(
        `[GRPC Server]: Getting all products in the category '${categoryName}'...`
      );

      if (!categoryName) {
        callback({
          code: grpc.status.INVALID_ARGUMENT,
          details: `The param 'name' was not provided`,
        });
      }

      var productsInCategory = products.filter(
        (x) => x.category === categoryName
      );

      callback(null, { products: productsInCategory });
    },

    getAllCategories: (_, callback) => {
      console.log("[GRPC Server]: Getting all categories...");

      var categories = [];

      products.forEach((x) => {
        var category = {
          name: x.category,
          thumbnail: x.thumbnail,
          image: x.images[0],
        };

        if (!categories.find((x) => x.name === category.name)) {
          categories.push(category);
        }
      });

      callback(null, { categories });
    },
    getProduct: (call, callback) => {
      var productId = call.request.id;

      console.log(
        `[GRPC Server]: Searching product with id '${productId}' ...`
      );

      if (!productId) {
        callback({
          code: grpc.status.INVALID_ARGUMENT,
          details: `The param 'id' was not provided`,
        });
      }

      let product = products.find((p) => p.id == productId);

      if (product) {
        callback(null, product);
      } else {
        callback({
          code: grpc.status.NOT_FOUND,
          details: `The product with id '${productId}' could not be found`,
        });
      }
    },
  });

  return server;
};

var server = getGrpcServer();

// launch the server and listen to the designated localhost port
server.bindAsync(
  `127.0.0.1:${PORT}`,
  grpc.ServerCredentials.createInsecure(),
  (error, port) => {
    // check if an error was noted when launching the server
    assert.ifError(error);

    console.log(
      `GRPC Server running at port ${port}: URL -> http:\\\\localhost:${port}`
    );

    server.start();
  }
);
