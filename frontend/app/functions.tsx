import {PathType} from '@/store/store'

export function isRegularPath(path: PathType | PathType[]): path is PathType {
	return !Array.isArray(path) && (path as PathType).Title !== undefined
}
