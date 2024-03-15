import { create } from 'zustand'
import { devtools, persist } from 'zustand/middleware' // для анализа в браузере
import instance from '../utils/axios'
import accessors from "../components/Dictionaries/Accessors/Accessors";

export const useAccessors = create(devtools((set, get) => ({
    accessors: [],
    loading: false,
    error: null,
    /*
    createAccessors: (title) => set(state => {
        const newAccessor = {
            id: new Date(),
            title: title
        }

        return {
            accessors: [...state.accessors, newAccessors]
        }
    })
     */
    /*
    createAccessor: (title) => set(state => ({
        accessors: [...state.accessors, {
            id: new Date(),
            title: title
        }]
    }))
     */
    createAccessor: (title) => {
        const newAccessor = {
            id: new Date(),
            title: title
        }

        set({accessors: [...get().accessors, newAccessor]})
    },

    //TODO: реализовать crud
    //Read
    readAccessors: async () => {
        set({ loading: true })
        try {
            const resp = await instance.get('api/v1/accessor/read')
            // console.log('this is data', resp.data)
            set({ accessors: resp.data, error: null })
        } catch (error) {
            set({ error: error.message })
        } finally {
            set({ loading: false })
        }
    },
    //Delete
    deleteAccessor: async (id) => {
        set({ loading: true })
        try {
            const resp = await instance.delete(`api/v1/accessor/delete/${id}`)
            const data = accessors.filter((accessor) => accessor.id !== id)
            set( { accessors: data, error: null })
        } catch (error) {
            set({ error: error.message })
        } finally {
            set({ loading: false })
        }
    }
})))

export const useJWT = create(devtools((set, get) => ({
    jwt: 'empty',
    setJWT: (token) => {
        set({ jwt: token })
    }
})))