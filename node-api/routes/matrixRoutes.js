const express = require("express");
const router = express.Router();
const { processMatrices } = require("../controllers/matrixController");

router.post("/process", processMatrices);

module.exports = router;
