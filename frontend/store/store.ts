import { create } from "zustand";


interface suggestionType {
    title: string;
    description: string;
    thumbnailUrl: string;
}

interface FormState {
    startSuggestions: suggestionType[];
    goalSuggestions: suggestionType[];
    openStartSuggestions: boolean;
    openGoalSuggestions: boolean;
    config: { start: string; goal: string; algorithm: string };
    setStartSuggestions: (startSuggestions: suggestionType[]) => void;
    setGoalSuggestions: (goalSuggestions: suggestionType[]) => void;
    setOpenStartSuggestions: (openStartSuggestions: boolean) => void;
    setOpenGoalSuggestions: (openGoalSuggestions: boolean) => void;
    setConfig: (config: { start: string; goal: string; algorithm: string }) => void;
}

interface ResultState {
    isLoading: boolean;
    result: { time: number; checkedArticles: number; passedArticles: number; path: string[] };
    setIsLoading: (isLoading: boolean) => void;
    setResult: (result: { time: number; checkedArticles: number; passedArticles: number; path: string[] }) => void;
}   



export const useResultStore = create<ResultState>((set) => ({
    isLoading: false,
    result: {time: 0, checkedArticles: 0, passedArticles: 0,path:[]},
    setIsLoading: (isLoading) => set({ isLoading }),
    setResult: (result) => set({ result }),
}));


export const useFormStore = create<FormState>((set) => ({
    startSuggestions: [],
    goalSuggestions: [],
    openStartSuggestions: false,
    openGoalSuggestions: false,
    config: { start: '', goal: '', algorithm: 'IDS' },
    setStartSuggestions: (startSuggestions) => set({ startSuggestions }),
    setGoalSuggestions: (goalSuggestions) => set({ goalSuggestions }),
    setOpenStartSuggestions: (openStartSuggestions) => set({ openStartSuggestions }),
    setOpenGoalSuggestions: (openGoalSuggestions) => set({ openGoalSuggestions }),
    setConfig: (config) => set({ config }),
}));