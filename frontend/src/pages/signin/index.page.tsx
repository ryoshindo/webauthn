import { useRouter } from "next/router";
import { FC } from "react";
import { Signin } from "./Signin";

const SigninPage: FC<{}> = () => {
  const router = useRouter();

  return <Signin onSubmit={() => {}} />;
};

const Page: FC<{}> = ({}) => {
  return <SigninPage />;
};

export default Page;
