export function maskCurrency(value: number): string {
  return value.toLocaleString("en-US", {
    style: "currency",
    currency: "USD",
  });
}

export function maskDate(date: string): string {
  return new Date(date).toLocaleDateString("en-US");
}

export function getUserLocale(): string {
  return navigator.language;
}

export function isUserPortugueseSpeaker(): boolean {
  return getUserLocale().startsWith("pt");
}
