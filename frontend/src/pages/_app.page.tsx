import "../../styles/globals.css";
import { ApolloProvider } from "@apollo/client";
import type { AppProps } from "next/app";
import { ChakraProvider } from "@chakra-ui/react";
import { client } from "src/apollo/client";

export default function App({ Component, pageProps }: AppProps) {
  const apolloClient = client;

  return (
    <ApolloProvider client={apolloClient}>
      <ChakraProvider>
        <Component {...pageProps} />
      </ChakraProvider>
    </ApolloProvider>
  );
}
