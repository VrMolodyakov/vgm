import React, { useCallback, useEffect, useState } from 'react'
import { useDependentState } from './hook/dependency'

export interface InputProps<T> {
  value: T
  onChange: (value: T) => void
  parseValue: (str: string) => Error | T
  asString: (value: T) => string
}

export function isError(value: unknown): value is Error {
    return value instanceof Error
  }

const ValidatedInput = <T extends Object>({
  value: inputValue,
  onChange,
  parseValue,
  asString
}: InputProps<T>) => {
  const [isValid, setIsValid] = useState<boolean>(true);

  const valueStr = asString(inputValue)
  const [value, setValue] = useDependentState(valueStr)

  const onChangeInternal = useCallback(function onChangeInternal(
    { currentTarget: { value } }: React.ChangeEvent<HTMLInputElement>
  ) {
    setValue(value)
    const parsed = parseValue(value)

    // if parsing fails, mark as invalid
    // and return
    if(isError(parsed)) {
      setIsValid(false)
      return
    }

    // parsing was successful
    setIsValid(true)
    onChange(parsed)
  }, [])

  return (
    <input
      onChange={onChangeInternal}
      style={{borderColor: isValid ? undefined : 'red'}}
      value={value}
    />
  )
}

export default ValidatedInput