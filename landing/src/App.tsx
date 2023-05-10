import { Token } from "@stripe/stripe-js";
import Checkout from "./components/Checkout";
import Signup from "./components/Signup";
import OTP from "./components/OTP";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Landing from "./components/Landing";

function App() {
  const router = createBrowserRouter([
    { path: "/", element: <Landing /> },
    { path: "/signup", element: <Signup /> },
    { path: "/verify-otp", element: <OTP /> },
    {
      path: "/checkout",
      element: (
        <Checkout
          handler={(token: Token) => {
            console.log(token);
          }}
        />
      ),
    },
  ]);
  return (
    <>
      <RouterProvider router={router} />
    </>
  );
}

export default App;
