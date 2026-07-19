package engine

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/share"
)

const (
	internalConnectionKeyAttrV1EngineParseDocument        = "_xaligoConnectionKey"
	internalConnectionSrcKeyAttrV1EngineParseDocument     = "_xaligoConnectionSrcKey"
	internalConnectionDstKeyAttrV1EngineParseDocument     = "_xaligoConnectionDstKey"
	internalConnectionSrcFrameAttrV1EngineParseDocument   = "_xaligoConnectionSrcFrame"
	internalConnectionDstFrameAttrV1EngineParseDocument   = "_xaligoConnectionDstFrame"
	internalConnectionCrossFrameAttrV1EngineParseDocument = "_xaligoConnectionCrossFrame"
)

var (
	IUPP001V1EngineParseDocument    = share.NewMCode("IUPP-001", "Parse read DSL failed")
	IUPP002V1EngineParseDocument    = share.NewMCode("IUPP-002", "Parse EOF branch")
	IUPP003V1EngineParseDocument    = share.NewMCode("IUPP-003", "Parse token failed")
	IUPP004V1EngineParseDocument    = share.NewMCode("IUPP-004", "Parse start element branch")
	IUPP005V1EngineParseDocument    = share.NewMCode("IUPP-005", "Parse item validation failed")
	IUPP006V1EngineParseDocument    = share.NewMCode("IUPP-006", "Parse connection validation failed")
	IUPP007V1EngineParseDocument    = share.NewMCode("IUPP-007", "Parse generic group validation failed")
	IUPP008V1EngineParseDocument    = share.NewMCode("IUPP-008", "Parse root assignment branch")
	IUPP009V1EngineParseDocument    = share.NewMCode("IUPP-009", "Parse child append branch")
	IUPP010V1EngineParseDocument    = share.NewMCode("IUPP-010", "Parse char data empty stack branch")
	IUPP011V1EngineParseDocument    = share.NewMCode("IUPP-011", "Parse text assignment branch")
	IUPP012V1EngineParseDocument    = share.NewMCode("IUPP-012", "Parse text append branch")
	IUPP013V1EngineParseDocument    = share.NewMCode("IUPP-013", "Parse unexpected closing tag branch")
	IUPP014V1EngineParseDocument    = share.NewMCode("IUPP-014", "Parse empty document branch")
	IUPP015V1EngineParseDocument    = share.NewMCode("IUPP-015", "Parse invalid root branch")
	IUPP016V1EngineParseDocument    = share.NewMCode("IUPP-016", "Parse expand connection shorthands failed")
	IUPP017V1EngineParseDocument    = share.NewMCode("IUPP-017", "Parse implicit V1 version branch")
	IUPP018V1EngineParseDocument    = share.NewMCode("IUPP-018", "Parse legacy V1 root branch")
	IUPVGGN001V1EngineParseDocument = share.NewMCode("IUPVGGN-001", "Validate generic group empty icon ID branch")
	IUPVGGN002V1EngineParseDocument = share.NewMCode("IUPVGGN-002", "Validate generic group invalid icon ID branch")
	IUPECS001V1EngineParseDocument  = share.NewMCode("IUPECS-001", "Expand connection shorthands item branch")
	IUPECS002V1EngineParseDocument  = share.NewMCode("IUPECS-002", "Expand connection shorthands item ID branch")
	IUPECS003V1EngineParseDocument  = share.NewMCode("IUPECS-003", "Expand connection shorthands empty alias branch")
	IUPECS004V1EngineParseDocument  = share.NewMCode("IUPECS-004", "Expand connection shorthands alias without ID branch")
	IUPECS005V1EngineParseDocument  = share.NewMCode("IUPECS-005", "Expand connection shorthands duplicate alias branch")
	IUPECS006V1EngineParseDocument  = share.NewMCode("IUPECS-006", "Expand connection shorthands collect failed")
	IUPECS007V1EngineParseDocument  = share.NewMCode("IUPECS-007", "Expand connection shorthands empty line branch")
	IUPECS008V1EngineParseDocument  = share.NewMCode("IUPECS-008", "Expand connection shorthands invalid shorthand branch")
	IUPECS009V1EngineParseDocument  = share.NewMCode("IUPECS-009", "Expand connection shorthands non shorthand branch")
	IUPECS010V1EngineParseDocument  = share.NewMCode("IUPECS-010", "Expand connection shorthands missing source branch")
	IUPECS011V1EngineParseDocument  = share.NewMCode("IUPECS-011", "Expand connection shorthands missing destination branch")
	IUPECS012V1EngineParseDocument  = share.NewMCode("IUPECS-012", "Expand connection shorthands traffic branch")
	IUPPA001V1EngineParseDocument   = share.NewMCode("IUPPA-001", "Position at negative offset branch")
	IUPPA002V1EngineParseDocument   = share.NewMCode("IUPPA-002", "Position at overflow offset branch")
	IUPVIN001V1EngineParseDocument  = share.NewMCode("IUPVIN-001", "Validate item spacer branch")
	IUPVIN002V1EngineParseDocument  = share.NewMCode("IUPVIN-002", "Validate item comma branch")
	IUPVIN003V1EngineParseDocument  = share.NewMCode("IUPVIN-003", "Validate item non numeric branch")
	IUPVCN001V1EngineParseDocument  = share.NewMCode("IUPVCN-001", "Validate connection missing source branch")
	IUPVCN002V1EngineParseDocument  = share.NewMCode("IUPVCN-002", "Validate connection missing destination branch")
	IUPVCN003V1EngineParseDocument  = share.NewMCode("IUPVCN-003", "Validate connection non numeric source branch")
	IUPVCN004V1EngineParseDocument  = share.NewMCode("IUPVCN-004", "Validate connection non numeric destination branch")
	IUPVCN005V1EngineParseDocument  = share.NewMCode("IUPVCN-005", "Validate connection nested branch")
	IUPVCN006V1EngineParseDocument  = share.NewMCode("IUPVCN-006", "Validate connection missing endpoint item branch")
	IUPVCN007V1EngineParseDocument  = share.NewMCode("IUPVCN-007", "Validate connection ambiguous endpoint item branch")
)

func ParseV1EngineParseDocument(r io.Reader) (entity.Document, error) {
	return ParseWithImportsV1EngineParseDocument(r, nil)
}

// ParseWithImportsV1EngineParseDocument parses a document and resolves explicitly
// supplied file imports before shared table normalization.
func ParseWithImportsV1EngineParseDocument(r io.Reader, imports *entity.ImportSource) (entity.Document, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		loggerV1EngineSharedLogging.ERROR(IUPP001V1EngineParseDocument, "read DSL failed", map[string]any{"error": err})
		return entity.Document{}, fmt.Errorf("read DSL: %w", err)
	}
	dec := xml.NewDecoder(bytes.NewReader(data))
	var stack []*entity.Node
	var root *entity.Node

	for {
		offset := int(dec.InputOffset())
		tok, err := dec.Token()
		if err == io.EOF {
			loggerV1EngineSharedLogging.DEBUG(IUPP002V1EngineParseDocument, "branch EOF")
			break
		}
		if err != nil {
			loggerV1EngineSharedLogging.ERROR(IUPP003V1EngineParseDocument, "token failed", map[string]any{"offset": offset, "error": err})
			return entity.Document{}, &entity.ParseError{Position: positionAtV1EngineParseDocument(data, offset), Err: fmt.Errorf("parse xml-like token: %w", err)}
		}

		switch t := tok.(type) {
		case xml.StartElement:
			loggerV1EngineSharedLogging.DEBUG(IUPP004V1EngineParseDocument, "branch start element", map[string]any{"tag": t.Name.Local})
			node := &entity.Node{Tag: t.Name.Local, Attrs: map[string]string{}, Position: positionAtV1EngineParseDocument(data, offset)}
			for _, a := range t.Attr {
				node.Attrs[a.Name.Local] = a.Value
			}
			if len(stack) > 0 {
				if err := validateNestedVersionV1EngineParseNode(node, stack[len(stack)-1]); err != nil {
					return entity.Document{}, &entity.ParseError{Position: node.Position, Err: err}
				}
			}
			if node.Tag == "item" {
				if err := validateItemNodeV1EngineParseNode(node); err != nil {
					loggerV1EngineSharedLogging.ERROR(IUPP005V1EngineParseDocument, "item validation failed", map[string]any{"error": err})
					return entity.Document{}, &entity.ParseError{Position: node.Position, Err: fmt.Errorf("parse <item>: %w", err)}
				}
			}
			if node.Tag == "connection" || node.Tag == "connections" {
				if err := validateConnectionNodeV1EngineParseConnection(node); err != nil {
					loggerV1EngineSharedLogging.ERROR(IUPP006V1EngineParseDocument, "connection validation failed", map[string]any{"error": err})
					return entity.Document{}, &entity.ParseError{Position: node.Position, Err: fmt.Errorf("parse <%s>: %w", node.Tag, err)}
				}
			}
			if node.Tag == "generic-group" {
				if err := validateGenericGroupNodeV1EngineParseNode(node); err != nil {
					loggerV1EngineSharedLogging.ERROR(IUPP007V1EngineParseDocument, "generic group validation failed", map[string]any{"error": err})
					return entity.Document{}, &entity.ParseError{Position: node.Position, Err: fmt.Errorf("parse <generic-group>: %w", err)}
				}
			}
			if isConnectableFrameTagV1EngineParseNode(node.Tag) && !isUMLCompartmentStartV1EngineParseDocument(stack, node) {
				if err := validateConnectableFrameNodeV1EngineParseNode(node); err != nil {
					loggerV1EngineSharedLogging.ERROR(IUPP007V1EngineParseDocument, "connectable frame validation failed", map[string]any{"error": err})
					return entity.Document{}, &entity.ParseError{Position: node.Position, Err: err}
				}
			}
			if len(stack) == 0 {
				loggerV1EngineSharedLogging.DEBUG(IUPP008V1EngineParseDocument, "branch root assignment", map[string]any{"tag": node.Tag})
				root = node
			} else {
				loggerV1EngineSharedLogging.DEBUG(IUPP009V1EngineParseDocument, "branch child append", map[string]any{"tag": node.Tag})
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, node)
			}
			stack = append(stack, node)
		case xml.CharData:
			if len(stack) == 0 {
				loggerV1EngineSharedLogging.DEBUG(IUPP010V1EngineParseDocument, "branch empty stack char data")
				continue
			}
			text := strings.TrimSpace(string(t))
			cur := stack[len(stack)-1]
			cur.TextRuns = append(cur.TextRuns, entity.TextRun{Text: string(t), Position: positionAtV1EngineParseDocument(data, offset)})
			if text != "" {
				if cur.Text == "" {
					loggerV1EngineSharedLogging.DEBUG(IUPP011V1EngineParseDocument, "branch text assignment", map[string]any{"tag": cur.Tag})
					cur.Text = text
				} else {
					loggerV1EngineSharedLogging.DEBUG(IUPP012V1EngineParseDocument, "branch text append", map[string]any{"tag": cur.Tag})
					cur.Text += " " + text
				}
			}
		case xml.EndElement:
			if len(stack) == 0 {
				loggerV1EngineSharedLogging.DEBUG(IUPP013V1EngineParseDocument, "branch unexpected closing tag", map[string]any{"tag": t.Name.Local})
				return entity.Document{}, &entity.ParseError{Position: positionAtV1EngineParseDocument(data, offset), Err: fmt.Errorf("unexpected closing tag: %s", t.Name.Local)}
			}
			stack = stack[:len(stack)-1]
		}
	}

	if root == nil {
		loggerV1EngineSharedLogging.ERROR(IUPP014V1EngineParseDocument, "branch empty document")
		return entity.Document{}, &entity.ParseError{Position: entity.Position{Line: 1, Column: 1}, Err: fmt.Errorf("empty document")}
	}
	if root.Tag != "xaligo" && root.Tag != "frame" && root.Tag != "frames" {
		loggerV1EngineSharedLogging.ERROR(IUPP015V1EngineParseDocument, "branch invalid root", map[string]any{"tag": root.Tag})
		return entity.Document{}, &entity.ParseError{Position: root.Position, Err: fmt.Errorf("root tag must be <xaligo>; legacy <frame> and <frames> roots are also accepted, got <%s>", root.Tag)}
	}
	if err := validateRootVersionV1EngineParseNode(root); err != nil {
		loggerV1EngineSharedLogging.ERROR(IUPP015V1EngineParseDocument, "branch unsupported root version", map[string]any{"tag": root.Tag, "version": root.Attrs["version"]})
		return entity.Document{}, &entity.ParseError{Position: root.Position, Err: err}
	}
	if _, specified := root.Attrs["version"]; !specified && root.Tag == "xaligo" {
		loggerV1EngineSharedLogging.WARN(IUPP017V1EngineParseDocument, implicitV1VersionWarningV1EngineParseNode(root), map[string]any{"tag": root.Tag})
	}
	envelope := root
	dataNode := (*entity.Node)(nil)
	legacyRoot := root.Tag != "xaligo"
	if legacyRoot {
		loggerV1EngineSharedLogging.WARN(IUPP018V1EngineParseDocument, legacyV1RootWarningV1EngineParseNode(root), map[string]any{"tag": root.Tag})
	} else {
		var err error
		root, dataNode, err = normalizeEnvelopeV1EngineParseNode(root)
		if err != nil {
			return entity.Document{}, err
		}
	}
	if err := normalizeUMLDiagramsV1EngineParseUml(root, dataNode); err != nil {
		return entity.Document{}, err
	}
	if err := validateFrameHierarchyV1EngineParseNode(root); err != nil {
		loggerV1EngineSharedLogging.ERROR(IUPP006V1EngineParseDocument, "frame hierarchy validation failed", map[string]any{"error": err})
		return entity.Document{}, err
	}
	if err := normalizeFrameMetadataV1EngineParseFrameMetadata(root, envelope); err != nil {
		return entity.Document{}, err
	}
	if err := resolveDatabaseImportsV1EngineParseDatabaseImport(dataNode, imports); err != nil {
		return entity.Document{}, err
	}
	if err := normalizeDatabasesV1EngineParseDatabase(root, dataNode); err != nil {
		return entity.Document{}, err
	}
	if err := resolveTableImportsV1EngineParseImport(root, dataNode, imports); err != nil {
		return entity.Document{}, err
	}
	if err := normalizeTablesV1EngineParseTable(root); err != nil {
		return entity.Document{}, err
	}
	assignConnectionKeysV1EngineParseNode(root)
	if err := expandConnectionShorthandsV1EngineParseShorthand(root, data); err != nil {
		loggerV1EngineSharedLogging.ERROR(IUPP016V1EngineParseDocument, "expand connection shorthands failed", map[string]any{"error": err})
		return entity.Document{}, err
	}
	if err := normalizeConnectionEndpointTagsV1EngineParseConnection(root); err != nil {
		loggerV1EngineSharedLogging.ERROR(IUPP006V1EngineParseDocument, "connection endpoint tag normalization failed", map[string]any{"error": err})
		return entity.Document{}, err
	}
	if err := validateEffectiveRouteArrowheadsV1EngineParseConnection(root); err != nil {
		loggerV1EngineSharedLogging.ERROR(IUPP006V1EngineParseDocument, "effective route arrowhead validation failed", map[string]any{"error": err})
		return entity.Document{}, err
	}
	if err := validateConnectionReferencesV1EngineParseReference(root); err != nil {
		loggerV1EngineSharedLogging.ERROR(IUPP006V1EngineParseDocument, "connection reference validation failed", map[string]any{"error": err})
		return entity.Document{}, err
	}

	return entity.Document{Root: root, Data: dataNode, Envelope: envelope, LegacyRoot: legacyRoot}, nil
}

func isUMLCompartmentStartV1EngineParseDocument(stack []*entity.Node, node *entity.Node) bool {
	if node == nil || len(stack) == 0 || !umlCompartmentTagsV1EngineParseUml[node.Tag] {
		return false
	}
	parent := stack[len(stack)-1]
	if parent == nil || !umlElementTagsV1EngineParseUml[parent.Tag] {
		return false
	}
	for index := len(stack) - 2; index >= 0; index-- {
		if stack[index].Tag == "uml" || stack[index].Tag == "uml-model" {
			return true
		}
	}
	return false
}

func positionAtV1EngineParseDocument(data []byte, offset int) entity.Position {
	if offset < 0 {
		loggerV1EngineSharedLogging.DEBUG(IUPPA001V1EngineParseDocument, "branch negative offset", map[string]any{"offset": offset})
		offset = 0
	}
	if offset > len(data) {
		loggerV1EngineSharedLogging.DEBUG(IUPPA002V1EngineParseDocument, "branch overflow offset", map[string]any{"offset": offset, "length": len(data)})
		offset = len(data)
	}
	prefix := data[:offset]
	line := bytes.Count(prefix, []byte{'\n'}) + 1
	lastNewline := bytes.LastIndexByte(prefix, '\n')
	column := offset - lastNewline
	return entity.Position{Offset: offset, Line: line, Column: column}
}
