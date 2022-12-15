import { useToast } from "@chakra-ui/react";
import { useRouter } from "next/router";
import { FC } from "react";

const RegisterPage: FC<{}> = () => {
  const router = useRouter();
  const toast = useToast();

  return <></>;
};

const Page: FC<{}> = ({}) => {
  return <RegisterPage />;
};

export default Page;
