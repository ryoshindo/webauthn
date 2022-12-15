import { useToast } from "@chakra-ui/react";
import { useRouter } from "next/router";
import { FC } from "react";
import { pagesPath } from "src/lib/$path";
import { useCreateAccountMutation } from "./document.gen";
import { Register } from "./Register";

const RegisterPage: FC<{}> = () => {
  const router = useRouter();
  const toast = useToast();

  const [createAccount, createAccountResult] = useCreateAccountMutation({
    onCompleted(data) {
      toast({
        title: "completed account registration",
        status: "success",
        position: "top",
      });

      router.push(pagesPath.register.webauthn.$url());
    },
  });

  return (
    <Register
      onSubmit={(values) => {
        createAccount({
          variables: {
            input: {
              email: values.email,
              userName: values.userName,
            },
          },
        });
      }}
      submitting={createAccountResult.loading}
    />
  );
};

const Page: FC<{}> = ({}) => {
  return <RegisterPage />;
};

export default Page;
