import { Link } from "react-router-dom";
import { useShop } from "../context/ShopContext";

function Navbar() {
  const { cart } = useShop();

  return (
    <nav className="navbar">
      <div className="navbar-inner">
        <div className="navbar-logo">Shop App</div>

        <div className="navbar-links">
          <Link to="/">Produkty </Link>
          <Link to="/cart">Koszyk ({cart.length}) </Link>
          <Link to="/payments">Płatności</Link>
        </div>
      </div>
    </nav>
  );
}

export default Navbar;