declare global {
  interface Window {
    currentUser?: { email: string; [key: string]: any };
  }
}
export {};