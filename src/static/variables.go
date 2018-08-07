// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package static

var LICENSE_NOTICE_TEMPLATE = map[string]string{
	"head": `<comment>  Copyright (c) <year> <author>`,
	"body-slc":
`
<comment>  Distributed under the <license name>,
<comment>  see the accompanying file LICENSE or <license link>

`,
	"body-mlc":
`
  Distributed under the <license name>,
  see the accompanying file LICENSE or <license link>
<comment>

`,
}
