import React from 'react'
import { dateString_MM_DD_YYYY, parseDate_MM_DD_YYYY } from './date'
import ValidatedInput from './validate'

interface DateInputProps {
  value: Date,
  onChange: (date: Date) => void
}

const DateInput: React.FC<DateInputProps> = ({
  value,
  onChange
}) => {
  return (
    <ValidatedInput
      onChange={onChange}
      value={value}
      parseValue={parseDate_MM_DD_YYYY}
      asString={dateString_MM_DD_YYYY}
    />
  )
}

export default DateInput