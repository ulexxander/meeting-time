import { ReactNode } from "react";

export const Page: React.FC<{ children: ReactNode }> = ({ children }) => {
  return <div>{children}</div>;
};
