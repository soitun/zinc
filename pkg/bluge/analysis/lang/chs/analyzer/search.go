/* Copyright 2022 Zinc Labs Inc. and Contributors
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package analyzer

import (
	"github.com/blugelabs/bluge/analysis"
	"github.com/go-ego/gse"

	"github.com/zincsearch/zincsearch/pkg/bluge/analysis/lang/chs/token"
	"github.com/zincsearch/zincsearch/pkg/bluge/analysis/lang/chs/tokenizer"
)

func NewSearchAnalyzer(seg *gse.Segmenter) *analysis.Analyzer {
	return &analysis.Analyzer{
		Tokenizer:    tokenizer.NewSearchTokenizer(seg),
		TokenFilters: []analysis.TokenFilter{token.NewStopTokenFilter(seg, nil)},
	}
}
