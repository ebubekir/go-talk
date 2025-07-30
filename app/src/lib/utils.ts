import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function getAvatarName(name: string) {
  if (!name) return "-"
  const parts = name.split(" ")
  if (parts.length === 1) {
    return parts[0].substring(0, 2).toUpperCase()
  }
  const firstInitial = parts[0].charAt(0).toUpperCase()
  const lastInitial = parts[parts.length - 1].charAt(0).toUpperCase()
  return `${firstInitial}${lastInitial}`
}
