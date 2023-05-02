// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate go run ../../generate/tags/main.go -ListTagsInIDElem=PipelineId -ServiceTagsSlice -TagOp=AddTags -TagInIDElem=PipelineId -UntagOp=RemoveTags -UpdateTags
// ONLY generate directives and package declaration! Do not add anything else to this file.

package datapipeline
