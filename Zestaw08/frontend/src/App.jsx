import { BrowserRouter, Routes, Route } from "react-router-dom";

import Login from "./Login";
import Register from "./Register";
import OAuthSuccess from "./OAuthSuccess";

function App() {
  return (
    <BrowserRouter>
      <Routes>

        <Route path="/" element={<Login />} />

        <Route path="/login" element={<Login />} />

        <Route path="/register" element={<Register />} />

        <Route
          path="/oauth-success"
          element={<OAuthSuccess />}
        />

      </Routes>
    </BrowserRouter>
  );
}

export default App;