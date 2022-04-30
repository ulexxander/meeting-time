import { ReactNode } from "react";

export const Page: React.FC<{ children: ReactNode }> = ({ children }) => {
  return <div className="mx-auto max-w-xl py-4">{children}</div>;
};
