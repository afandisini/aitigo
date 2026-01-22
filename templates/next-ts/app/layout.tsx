import type { ReactNode } from 'react';

export const metadata = {
  title: 'AitiGo Next.js Template',
  description: 'Next.js App Router + TypeScript + ESLint + Prettier',
};

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}