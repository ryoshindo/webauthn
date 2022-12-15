import { useRouter } from "next/router";
import { FC } from "react";
import { Signin } from "./SignIn";

const SigninPage: FC<{}> = () => {
  const router = useRouter();

  return <Signin />;
};

const Page: FC<{}> = ({}) => {
  return <SigninPage />;
};

export default Page;
