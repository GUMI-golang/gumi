package renderline

import "image"

type Worker interface {
	// 캐싱된 작업,
	// 대부분의 경우 이게 기본 옵션이고, 성능향상에 자동으로 기여한다.
	// 단 캐싱을 불가능하게 만들기 때문에 자신에게 할당된 영역 밖을 그리는 작업은 불가능하다.
	BaseRender(subimg *image.RGBA)
	// 캐싱된 작업,
	// 반드시 Base렌더링 이후 이뤄져야 하는 것들을 렌더링한다.
	// 성능에 악영향을 미칠 수 있기에 사용은 최소화하는게 좋다.
	// fullimg는 말 그대로 자신에게 주어진 영역 외 전체 이미지에 모두 영향을 미칠 수 있음을 의미한다.
	// updated는 업데이트 된 영역을 의미한다.
	DecalRender(fullimg *image.RGBA)
	DecalRenderRect() (updated image.Rectangle)
}