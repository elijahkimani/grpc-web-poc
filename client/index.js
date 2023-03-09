const client = require("./client");
const express = require("express");
const bodyParser = require("body-parser");

const PORT = process.env.PORT || 3000;

// create the express server
const app = express();

// setup http body parser middleware
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));

// add the routes
app.get("/", (req, res) => {
  client.getAll(null, (err, data) => {
    if (!err) {
      res.json(data);
    } else {
      res.send(err);
    }
  });
});

app.get("/products/:productId", (req, res) => {
  var productId = req.params.productId;
  client.get({ id: productId }, (err, data) => {
    if (!err) {
      res.json(data);
    } else {
      res.send(err);
    }
  });
});

// start the server
app.listen(PORT, () => {
  console.log("Server running at port %d", PORT);
});
