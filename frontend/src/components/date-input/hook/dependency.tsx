import { useEffect, useState } from "react"

export function useDependentState<T>(
    input: T
  ): [T, React.Dispatch<React.SetStateAction<T>>] {
    const [state, setState] = useState(input)
  
    useEffect(() => {
      setState(input)
    }, [input])
  
    return [state, setState]
  }