import { LoaderFunctionArgs } from "@remix-run/node";

export async function loader(args: LoaderFunctionArgs) {
    // provides data to the component
}

export default function RegisterPage() {
    return (
      <div>
        <h1>Login Page</h1>
        {/* Your login form and logic go here */}
      </div>
    );
  }