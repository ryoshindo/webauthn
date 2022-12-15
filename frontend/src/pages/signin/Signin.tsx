import { Box, Card, CardBody, ChakraProps, Input } from "@chakra-ui/react";
import { FC } from "react";

export const Signin: FC<{}> = ({}) => {
  return (
    <Box>
      <EmailForm />
    </Box>
  );
};

const EmailForm: FC<{} & ChakraProps> = ({ ...props }) => {
  return (
    <Box {...props}>
      <Card>
        <CardBody>
          <Box>Input your email.</Box>
          <Input />
        </CardBody>
      </Card>
    </Box>
  );
};
