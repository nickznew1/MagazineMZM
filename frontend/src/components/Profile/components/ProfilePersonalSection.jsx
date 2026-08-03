import {memo, useEffect, useState} from "react";
import {useForm} from "react-hook-form";
import "./style.css"
import {config} from "../../../config";


const ProfilePersonal = ({profile}) => {
    const {
        register,
        handleSubmit,
        formState: { errors },
        reset,
    } = useForm({
        mode: "onBlur",
        values:{
            company:"",
            first_name:"",
            second_name:"",
        }

    });


    const validateName = (value) => {
        if (!/^[а-яА-ЯёЁ]+$/.test(value)) {
            return 'Имя должно быть на русском языке';
        }
        return true
    }

    const validateLastName = (value) => {
        if (!/^[а-яА-ЯёЁ]+$/.test(value)) {
            return 'Фамилия должна быть на русском языке';
        }
        return true
    }


    const { user, setProfile} = profile;
    console.log(profile)


    useEffect(() => {
        if (user?.personal) {
            reset({
                company:user.user_personal.company || "",
                first_name:user.user_personal.first_name || "",
                second_name:user.user_personal.second_name || "",
            })
        }
    }, [user, reset]);

    const [visiblePersonal, setVisiblePersonal] = useState(false)
    const toggleVisiblePersonal = () => {
        setVisiblePersonal(!visiblePersonal);
    };

    const handleSubmitPersonal = async (data) => {
        const payload = {
              id: user.profileData.user_ordinary.id,
            ...data
        };
        try {
            const response = await fetch(`${config.apiUrl}/profile/personal`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(payload),
            })



            setProfile({
                ...user,
                profileData:{
                    ...user.profileData,
                    personal:payload
                }
            })
            setVisiblePersonal(false);
        } catch (error){
            throw new Error(error)
        }
    }

    const handleUpdatePersonal = async (data) => {
        const payload = {
            id: user.profileData.user_ordinary.id,
            ...data
        };
        try {
            const response = await fetch(`${config.apiUrl}/profile/personal/up`, {
                method: 'PATCH',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(payload),
            })

            const result = await response.json();

            setProfile({
                ...user,
                profileData:{
                    ...user.profileData,
                    user_personal:payload
                }
            })
            setVisiblePersonal(false);
        } catch (error){
            throw new Error(error)
        }
    }

    return (
        <>
        <h1>Ваши данные:</h1>
        <div className = "user-profile-personal-wrapper">
        {user.profileData.user_personal?.id ? (
                <>
                <div className = "user-profile-change__btn-wrapper">
                <div className = "user-profile-change__btn">
                    <button type ="button" className = {visiblePersonal ? "active" : ""} onClick = {toggleVisiblePersonal}>
                        Изменить персональные данные
                    </button>
                </div>
                </div>
                    {visiblePersonal ? (

                        <form onSubmit={handleSubmit(handleUpdatePersonal)}>
                            <div className="user-profile-change__inputs-wrapper">
                            <div className="user-profile-change__inputs">

                                <input type="text"
                                       placeholder = "Введите компанию"
                                       style={{ borderColor: errors.company ? 'red' : '#ccc'  }}
                                       {...register("company", {
                                       })
                                       }>
                                </input>
                                {errors.company && <span className = "user-profile-errors">{errors.company.message}</span> }
                            </div>
                            <div className="user-profile-change__inputs">
                                <input type="text"
                                       name="first_name"
                                       placeholder = "Введите ваше имя"
                                       style={{ borderColor: errors.first_name ? 'red' : '#ccc'  }}
                                       {...register("first_name", {
                                           validate: validateName,
                                       })
                                       }>
                                </input>
                                {errors.first_name && <span className = "user-profile-errors">{errors.first_name.message}</span> }
                            </div>
                            <div className="user-profile-change__inputs">
                            <input type="text"
                                   placeholder = "Введите вашу фамилию"
                                   style={{ borderColor: errors.second_name ? 'red' : '#ccc'  }}
                                   {...register("second_name", {
                                       validate: validateLastName,
                                   })
                                   }>
                            </input>
                            {errors.second_name && <span className = "user-profile-errors">{errors.second_name.message}</span> }
                            </div>

                            <div className="user-profile-submit__btn">
                                <button type="submit">Отправить</button>
                            </div>
                                </div>
                        </form>

                    ) : (
                        <div className="user-profile-personal-info">
                            <span>Ваше имя: {user.profileData.user_personal.first_name}</span>
                            <span>Ваша фамилия: {user.profileData.user_personal.second_name}</span>
                            <span>Ваша компания: {user.profileData.user_personal.company}</span>
                        </div>
                    )}
                </>
            ) : (

                <form onSubmit={handleSubmit(handleSubmitPersonal)}>
                    <div className="user-profile-change__inputs-wrapper">
                    <div className="user-profile-change__inputs">
                        <input type="text"
                               placeholder = "Введите компанию"
                               style={{ borderColor: errors.company ? 'red' : '#ccc'  }}
                               {...register("company", {
                               })
                               }>
                        </input>
                        {errors.company && <span className = "user-profile-errors">{errors.company.message}</span> }
                    </div>
                    <div className="user-profile-change__inputs">
                        <input type="text"
                               placeholder = "Введите ваше имя"
                               style={{ borderColor: errors.first_name ? 'red' : '#ccc'  }}
                               {...register("first_name", {
                                   validate: validateName,
                               })
                               }>
                        </input>
                        {errors.first_name && <span className = "user-profile-errors">{errors.first_name.message}</span> }
                    </div>
                    <div className="user-profile-change__inputs">
                    <input type="text"
                           placeholder = "Введите вашу фамилию"
                           style={{ borderColor: errors.second_name ? 'red' : '#ccc'  }}
                           {...register("second_name", {
                               validate: validateLastName,
                           })
                           }>
                    </input>
                    {errors.second_name && <span className = "user-profile-errors">{errors.second_name.message}</span> }
                    </div>
                    <div className="user-cab__accept-btn">
                        <button type="submit">Сохранить данные</button>
                    </div>
                    </div>
                </form>
            )}
        </div>
        </>
    )
}

export default memo(ProfilePersonal);