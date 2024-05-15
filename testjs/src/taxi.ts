export const calculateFare = (distance: number, waitingTime: number) => {
    return 4 * distance + waitingTime
  }
   
  export const roundDistance = (distance: number) => {
    return Math.ceil(distance * 2) / 2
  }
   
  export const roundWaitingTime = (waitingTime: number) => {
    return Math.ceil(waitingTime)
  }
   
  export const fare = (distance: number, waitingTime: number) => {
    return calculateFare(roundDistance(distance), roundWaitingTime(waitingTime))
  }

  export const calculateTotalFare = (distance: number, waitingTime: number) => {
    let totalFare = fare(distance, waitingTime)
    return isReachMinimumFare(totalFare) ? totalFare : 35
  }

  export const isReachMinimumFare = (fare: number) => {
    return fare > 35
  }