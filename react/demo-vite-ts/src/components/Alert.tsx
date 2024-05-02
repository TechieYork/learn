import { ReactNode } from "react";

interface AlertProps {
    children: ReactNode;
}

const Alert = ({ children }: AlertProps) => {
  return (
    <div className="alert alert-danger" role="alert">
      A simple danger alert—check it out! {children}
    </div>
  );
};

export default Alert;
