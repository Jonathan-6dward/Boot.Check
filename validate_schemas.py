import json
from pathlib import Path

from jsonschema import Draft202012Validator, RefResolver

ROOT = Path(__file__).parent


def load(path: Path):
    with path.open(encoding="utf-8") as handle:
        return json.load(handle)


def main():
    evidence_schema = load(ROOT / "scaffold/collector/evidence.schema.json")
    verdict_schema = load(ROOT / "scaffold/api/verdict.schema.json")
    validator = Draft202012Validator(evidence_schema, resolver=RefResolver.from_schema(evidence_schema))
    for fixture in sorted((ROOT / "scaffold/tests").glob("*.json")):
        data = load(fixture)
        errors = sorted(validator.iter_errors(data), key=lambda error: list(error.path))
        if errors:
            for error in errors:
                print(f"{fixture.name}: {'/'.join(map(str, error.path))}: {error.message}")
            raise SystemExit(1)
        print(f"evidence schema OK: {fixture.name}")
    verdict = {
        "verdict": "inconclusive",
        "confidence": 0.2,
        "headline_plain": "Os dados não bastam.",
        "plain_language_summary": "Fixture de contrato.",
        "five_w_two_h": {key: "não observado" for key in ["what", "who", "when", "where", "why", "how", "impact"]} | {"evidence_ids": []},
        "mitre": [],
        "supporting_evidence": [],
        "limitations": [],
        "recommended_next_steps": [],
        "technical_appendix": {"processes": [], "persistence": [], "network": [], "notes": []},
        "safety_notice": "Triagem indicativa; não é aconselhamento jurídico.",
    }
    verdict_errors = sorted(Draft202012Validator(verdict_schema).iter_errors(verdict), key=lambda error: list(error.path))
    if verdict_errors:
        for error in verdict_errors:
            print(f"verdict: {'/'.join(map(str, error.path))}: {error.message}")
        raise SystemExit(1)
    print("verdict schema OK")


if __name__ == "__main__":
    main()
