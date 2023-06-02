import Signup from "./components/Signup";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Landing from "./components/Landing";
import Cancel from "./components/Cancel";
import Intro from "./components/Intro";

function checkSession() {
  if (localStorage.getItem("token")) return true;
  return false;
}

function gotoApp() {
  if (checkSession()) {
    window.open(
      `${import.meta.env.VITE_APP_ADDRESS}/sign-in?token=${localStorage.getItem(
        "token"
      )}`,
      "_self"
    );
  }
}

function App() {
  const router = createBrowserRouter([
    { path: "/", element: <Landing /> },
    { path: "/intro", element: <Intro /> },
    { path: "/sign-up", element: <Signup /> },
    { path: "/cancel", element: <Cancel /> },
  ]);

  gotoApp();

  return (
    <div className="font-[Comfortaa]">
      <RouterProvider router={router} />
    </div>
  );
}

export default App;
