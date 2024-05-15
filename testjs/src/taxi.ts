// export const Taxi = () => 55

export function TukTuk(runningDistance: number, waitTime: number) {
    runningDistance = RoundNumber(runningDistance)
    waitTime = RoundNumber(waitTime)
    let calculate = (4 * runningDistance) + (1 * waitTime)
    return calculate
}

function RoundNumber(value: number) {
    value = value < 0 ? 0 : value
    const integerPart = Math.floor(value);
    const decimalPart = value - integerPart;

  if (decimalPart >= 0.1 && decimalPart <= 0.4) {
    return integerPart + 0.5;
  } else if (decimalPart >= 0.6 && decimalPart <= 0.9) {
    return integerPart + 1;
  } else {
    // For other cases, return the number itself or handle differently if needed
    return Math.round(value);
  }
}