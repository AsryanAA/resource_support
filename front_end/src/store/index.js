import { create } from 'zustand'
import { devtools, persist } from 'zustand/middleware' // для анализа в браузере
import instance from '../utils/axios'

export const useAccessors = create(devtools((set, get) => ({
    accessors: [{
        id: new Date,
        title: 'Example'
    }],
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
            const resp = await instance.get('accessor/read')
            // console.log('this is data', resp.data)
            set({ accessors: resp.data, error: null })
        } catch (error) {
            set({ error: error.message })
        } finally {
            set({ loading: false })
        }
    }
})))