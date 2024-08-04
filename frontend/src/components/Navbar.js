import React from "react";
import { Nav, NavLink, NavMenu } from "./NavbarElements";

const Navbar = () => {
    return (
        <Nav>
            <NavMenu>
                <NavLink to="/">Home</NavLink>
                <NavLink to="/proxy">Proxy</NavLink>
                <NavLink to="/domains">Domains</NavLink>
            </NavMenu>
        </Nav>
    );
}
export default Navbar;