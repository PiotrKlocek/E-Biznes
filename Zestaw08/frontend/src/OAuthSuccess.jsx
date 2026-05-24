import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

function OAuthSuccess() {

  const navigate = useNavigate();

  useEffect(() => {

    const params = new URLSearchParams(
      window.location.search
    );

    const token = params.get("token");

    if (token) {

      localStorage.setItem("token", token);

      navigate("/");

    } else {

      navigate("/login");

    }

  }, []);

  return (
    <div>
      <h1>Logowanie...</h1>
    </div>
  );
}

export default OAuthSuccess;