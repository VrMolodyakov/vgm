export function dateString_MM_DD_YYYY(date: Date): string {
  const [year, month, day] = date.toISOString()
    .substr(0, 10)
    .split('-')

  return [month, day, year].join('-')
}

export function parseDate_MM_DD_YYYY(str: string): Date | Error {
  try {
    const matchResult = str.match(/^\d{2}-\d{2}-\d{4}$/);
    if (matchResult !== null) {
      const date = new Date(matchResult[0]);
      return isNaN(date.valueOf()) ? new Error('') : date;
    } else {
      throw new Error('');
    }
  } catch (e) {
    return e as Error;
  }
}