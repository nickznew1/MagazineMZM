import {useState} from "react";
import AdminPageFilterUsers from "./AdminPageFilterUsers";
import AdminPageFilterApplications from "./AdminPageFilterApplications";
import AdminPageFilterCreateItem from "./AdminPageFilterCreateItem";
import {useNavigate} from "react-router";
import AdminPageFilterAllItems from "./AdminPageFilterAllItems";
import {config} from "../../../config";

const AdminPageFilter = () =>{

    const [loading, setLoading] = useState(true);
    const[error, setError] = useState(null);
    const [allUsers, setAllUsers] = useState([]);
    const [active, setActive] = useState(null);
    const navigate = useNavigate()


    const handleActiveButton =(id) =>{
        if (id === buttons.getUsers){
            handleGetUsers()
        }
        if (id === buttons.back) {
            navigate("/profile")
        }
        setActive(active === id ? null : id)
    }

    const buttons = {
        getUsers:0,
        openApplications:1,
        createItem:2,
        updateItem:3,
        allItems:4,
        back:5,
    }


    const handleGetUsers = async () => {
        setLoading(true);
        setError(null);
        try {
            const response = await fetch(`${config.apiUrl}/admin/users`, {
                method: 'GET',
                headers: {
                    "Content-Type": "application/json",
                }
            })
            if (response.ok) {
                const result = await response.json();
                setAllUsers(result);
                setLoading(false);
            }
        } catch (error) {
            setError(error);
            setLoading(false);
            setAllUsers(null)
        } finally{
            setLoading(false);
        }
    }



    return(
        <div className = "admin_page-filter-wrapper">
            <div className = "admin_page-filter-controls">
                <button className = {active === buttons.getUsers ? "active" : ""} onClick={()=>handleActiveButton(buttons.getUsers)}>Все пользователи</button>
                    <button className = {active === buttons.openApplications ? "active" : ""} onClick={()=>handleActiveButton(buttons.openApplications)}>Текущие заявки</button>
                       <button className = {active === buttons.createItem ? "active" : ""} onClick={()=>handleActiveButton(buttons.createItem)}>Добавить товар на сайт</button>
                <button className = {active === buttons.allItems ? "active" : ""} onClick={()=>handleActiveButton(buttons.allItems)}>Все товары/удаление товара</button>
                <button className = {active === buttons.back ? "active" : ""} onClick={()=>handleActiveButton(buttons.back)}>Вернуться в профиль</button>

            </div>
                {active === buttons.getUsers && (
                    <AdminPageFilterUsers
                        users = {allUsers}
                        loading = {loading}
                        />
                )}
            {active === buttons.openApplications && (
                <AdminPageFilterApplications
                    />
            )}
            {active === buttons.createItem && (
                <AdminPageFilterCreateItem />
            )}
            {active === buttons.allItems && (
                <AdminPageFilterAllItems />
            )}

        </div>
            )
}

export default AdminPageFilter;