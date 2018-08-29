const express = require("express");
const app = express();

app.route("/:group").get((req, res) => {
  var group = req.params["group"];
  res.send("Hi group: " + group);
});

app.listen(8080, () => {
  console.log("Server started!");
});
