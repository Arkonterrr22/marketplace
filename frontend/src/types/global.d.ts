declare global {
  interface Window {
    currentUser?: { email: string; id: string; name: string; company: string; [key: string]: any };
  }
}
export {};