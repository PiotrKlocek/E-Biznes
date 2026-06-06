import { Link } from "react-router-dom";
import { useShop } from "../context/ShopContext";

function Navbar() {
  const { cart } = useShop();
  return (
    <nav>
      <span className="logo">ShopApp</span>
      <Link to="/">Produkty</Link>
      <Link to="/cart">
        Koszyk
        {cart.length > 0 && <span className="badge">{cart.length}</span>}
      </Link>
      <Link to="/payments">Płatność</Link>
    </nav>
  );
}

export default Navbar;
