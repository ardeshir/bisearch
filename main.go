package main

func bisearch(array []int, target int, lowIdx int, highIdx int) int {
   if highIdx < lowIdx {
      return -1
   }

  mid := int((lowIdx + highIdx) / 2 )

  if array[mid] > target {
      return bisearch(array, target, lowIdx, min)
  } else if array[mid] < target {
      return bisearch(array, target, mid+1, highIdx)
  } else {
      return mid
  }
} // end of bisearch

func itsearch(arr []int, targ int, lowInx int, highInx) int {
    startInx := lowInx
    endInx := highInx
    var midl int
    for startInx <= endInx {
      midl = int(( startInx + endInx) / 2)
      if arr[midl] > targ {
         endInx = midl
      } else if arr[midl] < targ {
        startInx = midl
      } else {
         return midl
      }
    } // end of for

      return -1
} // end of itsearch
