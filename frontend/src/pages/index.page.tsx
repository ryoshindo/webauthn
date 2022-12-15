import { Box, Spinner } from "@chakra-ui/react";
import { useRouter } from "next/router";
import { FC, ReactNode } from "react";
import { Account } from "src/types/graphql.gen";
import { useFetchViewerQuery } from "./document.gen";

const IndexPage: FC<{ viewer: Account }> = () => {
  return <Box />;
};

const ViewerLoader: FC<{
  children: (props: { viewer: Account }) => ReactNode;
}> = ({ children }) => {
  const router = useRouter();

  const { loading, data } = useFetchViewerQuery({
    onError(e) {
      router.push("/signin");
    },
  });

  if (loading || !data?.viewer) {
    return <Spinner />;
  }

  return <>{children({ viewer: data.viewer })}</>;
};

const Page: FC<{}> = ({}) => {
  return <ViewerLoader>{(props) => <IndexPage {...props} />}</ViewerLoader>;
};

export default Page;
