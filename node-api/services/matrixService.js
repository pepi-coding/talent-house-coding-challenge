function flatten(matrix) {
  return matrix.flat();
}

function isDiagonal(matrix) {
  return matrix.every((row, i) =>
    row.every((val, j) => (i === j ? true : val === 0))
  );
}

exports.calculateStats = (Q, R) => {
  const allValues = flatten(Q).concat(flatten(R));

  const max = Math.max(...allValues);
  const min = Math.min(...allValues);
  const sum = allValues.reduce((acc, val) => acc + val, 0);
  const average = sum / allValues.length;

  return {
    max,
    min,
    average,
    sum,
    isQDiagonal: isDiagonal(Q),
    isRDiagonal: isDiagonal(R),
  };
};
