import {PathType, BonusPathType} from '@/store/store'

export function isBonusPath(
	path: BonusPathType[] | PathType[]
): path is BonusPathType[] {
	return (path as BonusPathType[])[0].paths !== undefined
}
