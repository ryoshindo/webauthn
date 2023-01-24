import {
  Box,
  Button,
  Card,
  CardBody,
  ChakraProps,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
} from "@chakra-ui/react";
import { FC } from "react";
import { useForm } from "react-hook-form";
import { reactHookFormDefaultOptions } from "src/form/reactHookFormDefaultOptions";

type FormValues = {
  email: string;
};

export const Signin: FC<
  { onSubmit: (values: FormValues) => void; submitting: boolean } & ChakraProps
> = ({ onSubmit, submitting, ...props }) => {
  return (
    <Box {...props}>
      <EmailForm onSubmit={onSubmit} submitting={submitting} />
    </Box>
  );
};

const EmailForm: FC<
  { onSubmit: (values: FormValues) => void; submitting: boolean } & ChakraProps
> = ({ onSubmit, submitting, ...props }) => {
  const {
    handleSubmit,
    register,
    formState: { errors },
  } = useForm<FormValues>({
    ...reactHookFormDefaultOptions,
  });

  return (
    <Box
      display="flex"
      justifyContent="center"
      alignItems="center"
      position="relative"
      h="100vh"
      {...props}
    >
      <form onSubmit={handleSubmit((v) => onSubmit(v))}>
        <Card w="300px" h="300px" alignItems="center" justifyContent="center">
          <CardBody w="75%" position="absolute">
            <FormControl id="email" isInvalid={!!errors.email}>
              <FormLabel mx={0}>email</FormLabel>
              <Input
                type="text"
                inputMode="email"
                placeholder="email"
                {...register("email", { required: "email is required." })}
              />
              <FormErrorMessage>{errors.email?.message}</FormErrorMessage>
            </FormControl>
            <Box
              mt="24px"
              display="flex"
              alignItems="center"
              justifyContent="center"
            >
              <Button w="100px" type="submit" isLoading={submitting}>
                Sign In
              </Button>
            </Box>
          </CardBody>
        </Card>
      </form>
    </Box>
  );
};
