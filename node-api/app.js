require("dotenv").config();
const express = require("express");
const matrixRoutes = require("./routes/matrixRoutes");

const app = express();

app.use(express.json());

const version = process.env.API_VERSION || "v1";
app.use(`/api/${version}/matrix`, matrixRoutes);

const PORT = process.env.PORT || 4000;
app.listen(PORT, () => {
  console.log(`Node.js API running on port ${PORT}`);
});
