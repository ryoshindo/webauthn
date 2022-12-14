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
  { onSubmit: (values: FormValues) => void } & ChakraProps
> = ({ onSubmit, ...props }) => {
  return (
    <Box {...props}>
      <EmailForm onSubmit={onSubmit} />
    </Box>
  );
};

const EmailForm: FC<
  { onSubmit: (values: FormValues) => void } & ChakraProps
> = ({ onSubmit, ...props }) => {
  const {
    handleSubmit,
    register,
    formState: { errors },
  } = useForm<FormValues>({
    ...reactHookFormDefaultOptions,
  });

  return (
    <Box display="flex" justifyContent="center" alignItems="center" {...props}>
      <form onSubmit={handleSubmit((v) => onSubmit(v))}>
        <Card>
          <CardBody>
            <FormControl id="email" isInvalid={!!errors.email}>
              <FormLabel>email</FormLabel>
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
              <Button w="100px" type="submit">
                Sign In
              </Button>
            </Box>
          </CardBody>
        </Card>
      </form>
    </Box>
  );
};
