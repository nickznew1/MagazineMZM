import {useState, useEffect, createContext, useContext} from 'react'
import {config} from "../config";
const AuthContext = createContext();


export const AuthProvider = ({ children }) => {
    const roles = {
        adminUser: "admin",
        ordinaryUser:"ordinary",
    }

    const [user, setUser] = useState(null);
    const [userRole, setUserRole] = useState(null);

    const [token, setToken] = useState(
        localStorage.getItem("token")
    );
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

  useEffect(()=>{
      if (token){
          localStorage.setItem("token", token);
      } else {
          localStorage.removeItem("token");
      }
  },[token])

    useEffect(()=>{
        const checkAuth = async () => {
            if (!token){
                setLoading(false)
                return
            }
            try {
                setLoading(true)
                setError(null);
                const response = await fetch(`${config.apiUrl}/user/`, {
                    headers: {
                        Authorization : `Bearer ${token}`
                    }
                } );
                if (response.ok) {
                    const userData = await response.json();
                    setUser(userData);
                    setLoading(false)
                    if (userData.user_role === roles.adminUser) {
                        setUserRole(roles.adminUser);
                    }else if (userData.user_role === roles.ordinaryUser) {
                        setUserRole(roles.ordinaryUser)
                    }
                } else {
                    setUser(null);
                    setToken(null)
                    setLoading(false)
                }
            } catch (err) {
                setUser(null);
                setToken(null)
                setError(err.message);
                setLoading(false)
            } finally {
                setLoading(false);
            }
        };
        checkAuth()
    }, [token]);


    const logout = async () => {
        await fetch (`${config.apiUrl}/auth/logout`,{ credentials: 'include' });
        localStorage.removeItem("token");
        setUser(null);
    }

    return (
        <AuthContext.Provider value ={{user, setUser,roles,userRole, token, setToken, logout, loading, error}}>
            {children}
        </AuthContext.Provider>
    )
}

export const useAuth =() => useContext(AuthContext);

