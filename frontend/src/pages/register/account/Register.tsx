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
  userName: string;
};

export const Register: FC<
  { submitting: boolean; onSubmit: (values: FormValues) => void } & ChakraProps
> = ({ submitting, onSubmit, ...props }) => {
  return (
    <Box>
      <EmailForm submitting={submitting} onSubmit={onSubmit} />
    </Box>
  );
};

const EmailForm: FC<
  { submitting: boolean; onSubmit: (values: FormValues) => void } & ChakraProps
> = ({ submitting, onSubmit, ...props }) => {
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
            <FormControl mt="24px" id="userName" isInvalid={!!errors.userName}>
              <FormLabel>user name</FormLabel>
              <Input
                type="text"
                inputMode="text"
                placeholder="user name"
                {...register("userName", {
                  required: "user name is required.",
                })}
              />
              <FormErrorMessage>{errors.userName?.message}</FormErrorMessage>
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
