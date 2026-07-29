import {useEffect, useState} from "react";
import Modal from "react-modal";
import {config} from "../../../config";


const AdminPageFilterAllItems = () => {

    const [allItems, setAllItems] = useState([]);
    const [modal, setModal] = useState(false);

    const handleDeleteItem = async (id) => {
        try{
            const response = await fetch (`${config.apiUrl}/item/delete`, {
                method: "DELETE",
                headers: {
                    'Content-type':'application/json',
                },
                body:JSON.stringify({id:id})
            })

            if (response.ok) {
                setAllItems(allItems.filter((item)=>item.id !==id))
            }

        }catch(err){
            throw new Error(err)
        }
    }



    const openModal = (id) => {
        setModal(id)
    }
    const closeModal = () => {
        setModal(false)
    }

    useEffect(()=>{
        const handleGetAllItems = async () => {
            try {
                const response = await fetch (`${config.apiUrl}/item/all`, {
                    method: 'GET',
                    headers: {
                        'Content-Type' :'application/json'
                    },
                });
                if (response.ok){
                    const result = await response.json();
                    setAllItems(result)
                }
            } catch (err) {
                throw new Error(err)
            } finally {

            }
        }
        handleGetAllItems()

        },
        [])

    return (
        <div>
        <table className="admin_page-users-table">
            <thead>
            <tr>
                <th>ID</th>
                <th>Название</th>
                <th>Цена</th>
                <th>Тип</th>
                <th>Способ измерения</th>
                <th>Удалить</th>
            </tr>
            </thead>
        </table>
    <div className = "admin_page-table-body">
        <table className="admin_page-users-table">
            <tbody>
            {allItems && allItems.map((item) => (
                <tr key={item.id}>
                    <td>{item.id}</td>
                    <td>{item.name}</td>
                    <td>{item.price}</td>
                    <td>{item.item_type}</td>
                    <td>{item.item_secondary_type}</td>
                    <td><button onClick={()=>openModal(item.id)}>Удалить товар</button></td>
                    {modal === item.id && (
                        <Modal
                            isOpen={openModal}
                            onRequestClose={closeModal}
                            className = "admin-page-modal-content"
                            overlayClassName = "admin-page-modal-overlay" >
                            <h1>Вы действительно желаете удалить данный товар?</h1>
                            <p>{item.id}</p>
                            <p>{item.name}</p>
                            <p>{item.price}</p>
                            <p>{item.item_type}</p>
                            <p>{item.item_secondary_type}</p>
                            <button onClick = {()=>handleDeleteItem(item.id)}>Удалить</button>
                            <button onClick = {closeModal}>Выйти</button>

                        </Modal>
                    )}
                </tr>
            ))}
            </tbody>
        </table>
    </div>
    </div>
    )

}

export default AdminPageFilterAllItems