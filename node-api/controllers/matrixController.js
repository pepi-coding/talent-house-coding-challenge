const matrixService = require("../services/matrixService");

exports.processMatrices = async (req, res) => {
  try {
    const { Q, R } = req.body;

    if (!Q || !R) {
      return res.status(400).json({ error: "Q and R matrices are required" });
    }

    const result = matrixService.calculateStats(Q, R);
    res.json(result);
  } catch (err) {
    console.error("Error processing matrices:", err);
    res.status(500).json({ error: "Internal server error" });
  }
};
