import * as WebAuthnJSON from "@github/webauthn-json";
import { Spinner, useToast } from "@chakra-ui/react";
import { useRouter } from "next/router";
import { FC, ReactNode } from "react";
import {
  useCompleteWebauthnRegistrationMutation,
  useInitiateWebauthnRegistrationMutation,
} from "./document.gen";
import { Register } from "../account/Register";
import { Account } from "src/types/graphql.gen";
import { useFetchViewerQuery } from "src/pages/document.gen";

const RegisterPage: FC<{ viewer: Account }> = ({ viewer }) => {
  const router = useRouter();
  const toast = useToast();

  const [initiateWebauthnRegistration, initiateWebauthnRegistrationResult] =
    useInitiateWebauthnRegistrationMutation({
      onCompleted(data) {
        const options = JSON.parse(data.initiateWebauthnRegistration);
        WebAuthnJSON.create({ publicKey: options }).then((credential) => {
          completeWebauthnRegistration({
            variables: {
              input: {
                credential: JSON.stringify(credential),
              },
            },
          });
        });
      },
    });

  const [completeWebauthnRegistration, completeWebauthnRegistrationResult] =
    useCompleteWebauthnRegistrationMutation({
      onCompleted(data) {
        toast({
          title: "completed webauthn device registration",
          status: "success",
          position: "top",
        });
      },
    });

  return (
    <Register
      submitting={
        initiateWebauthnRegistrationResult.loading ||
        completeWebauthnRegistrationResult.loading
      }
      onSubmit={() => {
        initiateWebauthnRegistration({});
      }}
    />
  );
};

const ViewerLoader: FC<{
  children: (props: { viewer: Account }) => ReactNode;
}> = ({ children }) => {
  const { loading, data } = useFetchViewerQuery({});

  if (loading || !data?.viewer) {
    return <Spinner />;
  }

  return <>{children({ viewer: data.viewer })}</>;
};

const Page: FC<{}> = ({}) => {
  return <ViewerLoader>{(props) => <RegisterPage {...props} />}</ViewerLoader>;
};

export default Page;
