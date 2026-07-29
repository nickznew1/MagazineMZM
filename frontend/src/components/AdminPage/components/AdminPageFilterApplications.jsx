import Modal from "react-modal";
import {AnimatePresence, motion} from "framer-motion";
import {useEffect, useState} from "react";
import {config} from "../../../config";


const AdminPageFilterApplications = ({/*applications, setAllApplications*/}) =>{

    const [modalIsOpen, setModalIsOpen] = useState(false);
    const [application, setApplication] = useState(null);
    const [activeSpec, setActiveSpec] = useState(false)
    const [activeStatus, setActiveStatus] = useState(false);

    const[error, setError] = useState(null);


    const [allApplications, setAllApplications] = useState([]);
    const [filteredApplications, setFilteredApplications] = useState(false);

  useEffect(() => {
      const handleGetApplicationsActive = async () => {
          setError(null);
          try {
              const response = await fetch(`http://localhost:8080/admin/applications`, {
                  method: 'GET',
                  headers: {
                      "Content-Type": "application/json",
                  }
              })
              if (response.ok) {
                  const result = await response.json();
                  setAllApplications(result);
                  setError(null);
              }
          } catch (err) {
              setError(error);
              setAllApplications([])
          }
      }
      handleGetApplicationsActive()
  }, []);


    const handleSpec =(id) =>{
        setActiveSpec(activeSpec === id ? null : id)
    }
    const openModal = (id) => {
        setModalIsOpen(true);
        handleGetApplication(id)
    }

    const closeModal = () => {
        setModalIsOpen(false);
    }

    const handleFilterApplications =() => {
        if (!filteredApplications) {
            setAllApplications(applications =>
                applications.sort((a, b) => new Date(b.order_date) - new Date(a.order_date)) ? [...applications] : applications)
            setFilteredApplications(true)
        } else {
            setAllApplications(applications =>
                applications.sort((a, b) => new Date(a.order_date) - new Date(b.order_date)) ? [...applications] : applications)
            setFilteredApplications(false)
        }
    }

    const statusVariants = {
        0:"В обработке",
        1:"Подтвержден",
        2:"Отменен",
        3:"Отправлен заказчику",
        4:"Доставлен",
    }

    const handleStatusButton = (status, id) => {
        handleChangeStatus(status,id)
    }

    const handleActiveStatusButton = (id) =>{
        setActiveStatus(activeStatus === id ? null : id)
    }



    const handleGetApplication = async(id) =>{
        setError(null);
        try {
            const response = await fetch(`${config.apiUrl}/admin/application/${id}`, {
                method: 'GET',
                headers: {
                    "Content-Type": "application/json",
                },
            })
            if (response.ok) {
                const result = await response.json();
                setApplication(result)
                setError(null);
            }
        } catch(err){
            setError(err);
            setApplication(null)
        }
    }

    const handleChangeStatus = async (status,id) =>{

        setError(null);
        const payload = {
            application_status:status,
            id:id,
        }
        try {
            const response = await fetch(`${config.apiUrl}/admin/status`, {
                method: 'POST',
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(payload),
            })
            if (response.ok) {
                const result = await response.json();
                console.log(result)
                setActiveStatus(null)
                if (response.ok) {
                    setAllApplications(prev =>
                        prev.map(app =>
                            app.id === id
                                ? {...app, application_status: status}
                                : app
                        )
                    );
                }

                setError(null);
            }
        }catch(error){
            setAllApplications(null)
            throw new Error(error)

        }finally{

        }
    }

    return(
        <div>
        <table className="admin_page-users-table">
            <thead>
            <tr>
                <th>ID</th>
                <th>Пользователь</th>
                <th>Товары</th>
                <th><button id = {filteredApplications ? "application_filter_active" : "application_filter_disabled"} onClick = {()=>handleFilterApplications(!filteredApplications)}>

                    Дата заявки <svg className = {filteredApplications ? "active" : ""} viewBox="0 0 24 24" width="16" height="16" stroke="currentColor"
                                     strokeWidth="2" fill="none" strokeLinecap="round" strokeLinejoin="round">
                    <polyline points="6 9 12 15 18 9"></polyline>
                </svg>
                </button>
                </th>
                <th>Статус</th>
                <th>Изменить статус</th>
                <th>Посмотреть заявку</th>
            </tr>
            </thead>
        </table>
        <div className = "admin_page-table-body">
           <table className="admin_page-users-table">
            <tbody>
            {allApplications?.map((item) => (
                <tr key={item.id}>
                    <td>{item.id}</td>
                    <td>{item.login}</td>
                    <td className = "admin_application_name">
                        {Object.entries(item.items ?? {}).map(([key, value]) => (
                            <div key ={item.id}>
                                {value.name}
                            </div>
                        ))}
                    </td>
                    <td>{new Date(item.order_date)
                        .toLocaleDateString("ru-RU")}</td>
                    <td>
                        {item.application_status}
                    </td>
                    <td>
                        <div className="admin-page-item-status__buttons">
                            <button onClick ={()=>handleActiveStatusButton(item.id)}>Выберите статус</button>
                            {activeStatus === item.id &&
                                Object.values(statusVariants).filter((active) => active !== item.application_status).map((status) => (
                                    <button onClick = {()=>handleStatusButton(status, item.id)}>{status}</button>
                                ))
                            }
                        </div>
                    </td>
                    <td>
                        <button onClick = {()=>openModal(item.id)}>Посмотреть заявку</button>
                        <Modal
                            isOpen={modalIsOpen}
                            onRequestClose={closeModal}
                            className = "admin-page-modal-content"
                            overlayClassName = "admin-page-modal-overlay">
                            {application && application.items && application.items.length > 0 ? (
                                <div className ="admin-page-application-items-wrapper">
                                    {application.items.map((item, index) => (
                                        <div id = {index} className = "admin-page-application-item">
                                            <img alt = {item.name} src = {`${config.imgServerUrl}${item.item_picture}`}></img>
                                            <span id ="name">{item.name}</span>
                                            <span id ="item-type">Тип: {item.item_type}</span>
                                            <span id = "item-secondary_type">Тип измерения: {item.item_secondary_type}</span>
                                            <span id = "item-count"> Количество: {item.count}</span>
                                            <button className = {activeSpec === item.item_spec_id ? "active" : ""} onClick = {()=>handleSpec(item.item_spec_id)}>Показать выбранную спецификацию
                                                <svg className = "filter-dropdown-icon" viewBox="0 0 24 24" width="16" height="16" stroke="currentColor"
                                                     strokeWidth="2" fill="none" strokeLinecap="round" strokeLinejoin="round">
                                                    <polyline points="6 9 12 15 18 9"></polyline>
                                                </svg>
                                            </button>
                                            <AnimatePresence>
                                                {activeSpec === item.item_spec_id &&
                                                    <motion.div
                                                        initial = {{ height:0, opacity:0}}
                                                        animate ={{height:"auto", opacity:1}}
                                                        exit={{height:0, opacity:0}}
                                                        transition={{duration: 0.3}}
                                                        className = "admin-page-application-item__specs-wrapper">
                                                        <div className = "admin-page-application-item__specs">
                                                            {Object.entries(item.props).map(([key,value]) => (
                                                                <>
                                                                    <p key ={key} id ="key">{key}:</p>
                                                                    <p key ={value} id ="value">{value}</p>
                                                                </>
                                                            ))}
                                                        </div>
                                                    </motion.div>
                                                }
                                            </AnimatePresence>
                                        </div>

                                    ))}
                                </div>
                            ) : (
                                <p>Ошибка загрузки товара</p>
                            )}
                            {application && application.items && application.items.length > 0 &&
                                <div className="admin-page-application-item-info-wrapper">
                                    <div className="admin-page-application-item-info">
                                        <span className = "header">Имя заказчика: </span>
                                        <span>{application.first_name}</span>
                                    </div>
                                    <div className="admin-page-application-item-info">
                                        <span className = "header">Фамилия заказчика: </span>
                                        <span>{application.second_name}</span>
                                    </div>
                                    <div className="admin-page-application-item-info">
                                        <span className = "header">Город: </span>
                                        <span>{application.city}</span>
                                    </div>
                                    <div className="admin-page-application-item-info">
                                        <span className = "header">Название компании: </span>
                                        <span>{application.company}</span>
                                    </div>
                                    <div className="admin-page-application-item-info">
                                        <span className = "header">Адрес: </span>
                                        <span>{application.address}</span>
                                    </div>
                                    <div className="admin-page-application-item-info">
                                        <span className = "header">Email: </span>
                                        <span>{application.email}</span>
                                    </div>
                                    <div className="admin-page-application-item-info">
                                        <span className = "header">Телефон: </span>
                                        <span>{application.phone_number}</span>
                                    </div>
                                </div>
                            }
                            <button id = "close" onClick={closeModal}>Закрыть</button>
                        </Modal>
                    </td>
                </tr>
            ))}
            </tbody>

        </table>
        </div>
        </div>
    )
}

export default AdminPageFilterApplications;