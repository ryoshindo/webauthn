import * as WebAuthnJSON from "@github/webauthn-json";
import { useToast } from "@chakra-ui/react";
import { useRouter } from "next/router";
import { FC, useState } from "react";
import {
  useCompleteWebauthnLoginMutation,
  useInitiateWebauthnLoginMutation,
} from "./document.gen";
import { Signin } from "./Signin";

const SigninPage: FC<{}> = () => {
  const router = useRouter();
  const toast = useToast();

  const [email, setEmail] = useState<string>("");

  const [initiateWebauthnLogin, initiateWebauthnLoginResult] =
    useInitiateWebauthnLoginMutation({
      onCompleted(data) {
        const options = JSON.parse(data.initiateWebauthnLogin);
        console.log(options);
        WebAuthnJSON.get({ publicKey: options["publicKey"] }).then(
          (credential) => {
            completeWebauthnLogin({
              variables: {
                input: {
                  email: email,
                  credential: JSON.stringify(credential),
                },
              },
            });
          }
        );
      },
    });

  const [completeWebauthnLogin, completeWebauthnLoginResult] =
    useCompleteWebauthnLoginMutation({
      onCompleted(data) {
        toast({
          title: "success signin",
          status: "success",
          position: "top",
        });
      },
    });

  return (
    <Signin
      onSubmit={(values) => {
        setEmail(values.email);
        initiateWebauthnLogin({
          variables: {
            input: {
              email: values.email,
            },
          },
        });
      }}
      submitting={
        initiateWebauthnLoginResult.loading ||
        completeWebauthnLoginResult.loading
      }
    />
  );
};

const Page: FC<{}> = ({}) => {
  return <SigninPage />;
};

export default Page;
