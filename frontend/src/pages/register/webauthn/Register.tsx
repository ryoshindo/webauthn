import {
  Box,
  Button,
  Card,
  CardBody,
  ChakraProps,
  FormControl,
  FormLabel,
} from "@chakra-ui/react";
import { FC } from "react";
import { useForm } from "react-hook-form";
import { reactHookFormDefaultOptions } from "src/form/reactHookFormDefaultOptions";

type FormValues = {};

export const Register: FC<
  { onSubmit: (values: FormValues) => void; submitting: boolean } & ChakraProps
> = ({ submitting, onSubmit, ...props }) => {
  const {
    handleSubmit,
    formState: { errors },
  } = useForm<FormValues>({
    ...reactHookFormDefaultOptions,
  });

  return (
    <Box display="flex" justifyContent="center" alignItems="center" {...props}>
      <form onSubmit={handleSubmit((v) => onSubmit(v))}>
        <Card>
          <CardBody>
            <FormControl id="webauthn">
              <FormLabel>WebAuthn Device</FormLabel>
            </FormControl>
            <Box
              mt="24px"
              display="flex"
              alignItems="center"
              justifyContent="center"
            >
              <Button w="100px" type="submit" isLoading={submitting}>
                Register
              </Button>
            </Box>
          </CardBody>
        </Card>
      </form>
    </Box>
  );
};
